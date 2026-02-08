const Sphere = @import("sphere.zig").Sphere;
const Ray = @import("ray.zig").Ray;
const v = @import("vector.zig");
const std = @import("std");
const constants = @import("constants.zig");
const quadratics = @import("quadratics.zig");
const mathf = @import("mathf.zig");

pub const MAX_INTERSECTIONS = 4;

pub const Shape = union(enum) {
    none: void,
    sphere: Sphere,

    pub fn approxEq(self: Shape, other: Shape) bool {
        if (std.meta.activeTag(self) != std.meta.activeTag(other)) {
            return false;
        }

        return switch (self) {
            .sphere => |s| s.approxEq(other.sphere),
            .none => true,
        };
    }
};

pub const Intersection = struct {
    t: f32,
    object: Shape,

    pub fn init(t: f32, object: Shape) Intersection {
        return .{ .t = t, .object = object };
    }

    pub fn approxEq(self: Intersection, other: Intersection) bool {
        if (@abs(self.t - other.t) > constants.epsilon) {
            return false;
        }

        if (!self.object.approxEq(other.object)) {
            return false;
        }
        return true;
    }

    pub const empty = Intersection{
        .t = 0,
        .object = Shape{ .none = {} },
    };
};

pub const Intersections = struct {
    items: [MAX_INTERSECTIONS]Intersection,
    count: usize,

    pub fn get(self: Intersections, idx: usize) Intersection {
        return self.items[idx];
    }

    pub fn init(values: anytype) Intersections {
        const T = @TypeOf(values);
        const vals_type_info = @typeInfo(T);
        if (vals_type_info != .@"struct") {
            @compileError("expected tuple or struct argument, found " ++ @typeName(T));
        }

        const fields_info = vals_type_info.@"struct".fields;
        if (fields_info.len > MAX_INTERSECTIONS) {
            @compileError("amount of values:" ++ fields_info.len ++ " exceeded MAX_INTERSECTIONS: " ++ MAX_INTERSECTIONS);
        }

        var result = Intersections{
            .items = undefined,
            .count = fields_info.len,
        };

        // Use inline for with the actual fields length, then fill the rest
        inline for (fields_info, 0..) |_, i| {
            result.items[i] = values[i];
        }

        inline for (fields_info.len..MAX_INTERSECTIONS) |i| {
            result.items[i] = Intersection.empty;
        }

        if (result.count > 1) {
            std.mem.sort(Intersection, result.items[0..result.count], {}, struct {
                fn lessThan(_: void, a: Intersection, b: Intersection) bool {
                    return a.t < b.t; // Assuming you want to sort by 't' field
                }
            }.lessThan);
        }

        return result;
    }

    pub fn zero() Intersections {
        return init(.{});
    }

    pub fn one(t: Intersection) Intersections {
        return init(.{t});
    }

    pub fn two(t1: Intersection, t2: Intersection) Intersections {
        return init(.{ t1, t2 });
    }

    pub fn approxEq(self: Intersections, other: Intersections) bool {
        if (self.count != other.count) {
            return false;
        }

        for (0..self.count) |i| {
            const a = self.items[i];
            const b = other.items[i];

            if (@abs(a.t - b.t) > constants.epsilon) {
                return false;
            }

            if (!a.object.approxEq(b.object)) {
                return false;
            }
        }

        return true;
    }

    pub fn hit(self: Intersections) Intersection {
        // first non-negative intersection is the hit
        for (0..self.count) |i| {
            if (self.items[i].t >= 0) {
                return self.items[i];
            }
        }

        // we get here if there are no non-negative t values or of count is 0
        return Intersection.empty;
    }

    test "Intersections:hit when all intersections have positive t" {
        const s = Shape{ .sphere = Sphere.unit };
        const xs = Intersections.init(.{
            Intersection.init(2, s),
            Intersection.init(1, s),
        });
        const expected = Intersection.init(1, s);
        try std.testing.expect(xs.hit().approxEq(expected));
    }

    test "Intersections:hit when some intersections have negative t" {
        const s = Shape{ .sphere = Sphere.unit };
        const xs = Intersections.init(.{
            Intersection.init(1, s),
            Intersection.init(-1, s),
        });
        const expected = Intersection.init(1, s);
        try std.testing.expect(xs.hit().approxEq(expected));
    }

    test "Intersections:hit when all intersections have negative t" {
        const s = Shape{ .sphere = Sphere.unit };
        const xs = Intersections.init(.{
            Intersection.init(-41, s),
            Intersection.init(-1, s),
        });
        const expected = Intersection.empty;
        try std.testing.expect(xs.hit().approxEq(expected));
    }

    test "Intersections:hit is always the lowest non-negative intersection" {
        const s = Shape{ .sphere = Sphere.unit };
        const xs = Intersections.init(.{
            Intersection.init(5, s),
            Intersection.init(7, s),
            Intersection.init(-3, s),
            Intersection.init(2, s),
        });
        const expected = Intersection.init(2, s);
        try std.testing.expect(xs.hit().approxEq(expected));
    }
};

pub fn raySphere(ray: Ray, sphere: Sphere) Intersections {
    const D = ray.direction;
    const S = ray.origin;
    const C = sphere.center;
    const r = sphere.radius;

    const a = v.dot(D, D);
    const b = 2 * v.dot(S, D) - 2 * v.dot(D, C);
    const c = v.dot(S, S) - 2 * v.dot(S, C) + v.dot(C, C) - r * r;

    const shape = Shape{ .sphere = sphere };
    return switch (quadratics.solve(a, b, c)) {
        .Zero => Intersections.zero(),
        .One => |s| Intersections.one(Intersection.init(s, shape)),
        .Two => |s| {
            const t1 = @min(s[0], s[1]);
            const t2 = @max(s[0], s[1]);
            return Intersections.two(Intersection.init(t1, shape), Intersection.init(t2, shape));
        },
    };
}

test "intersect:raySphere when ray passes through two points on a sphere" {
    const r = Ray.init(v.point(0, 0, -5), v.vector(0, 0, 1));
    const s = Sphere.unit;
    const expected = Intersections.two(
        Intersection.init(4, Shape{ .sphere = s }),
        Intersection.init(6, Shape{ .sphere = s }),
    );
    try std.testing.expect(Intersections.approxEq(raySphere(r, s), expected));
}

test "intersect:raySphere when ray passes through one point on a sphere" {
    const r = Ray.init(v.point(0, 1, -5), v.vector(0, 0, 1));
    const s = Sphere.unit;
    const expected = Intersections.one(
        Intersection.init(5, Shape{ .sphere = s }),
    );
    try std.testing.expect(Intersections.approxEq(raySphere(r, s), expected));
}

test "intersect:raySphere when ray misses the sphere" {
    const r = Ray.init(v.point(0, 2, -5), v.vector(0, 0, 1));
    const s = Sphere.unit;
    const expected = Intersections.zero();
    try std.testing.expect(Intersections.approxEq(raySphere(r, s), expected));
}

test "intersect:raySphere when ray originates in the sphere" {
    const r = Ray.init(v.point(0, 0, 0), v.vector(0, 0, 1));
    const s = Sphere.unit;
    const expected = Intersections.two(
        Intersection.init(-1, Shape{ .sphere = s }),
        Intersection.init(1, Shape{ .sphere = s }),
    );
    try std.testing.expect(Intersections.approxEq(raySphere(r, s), expected));
}

test "intersect:raySphere when sphere is behind the ray" {
    const r = Ray.init(v.point(0, 0, 5), v.vector(0, 0, 1));
    const s = Sphere.unit;
    const expected = Intersections.two(
        Intersection.init(-6, Shape{ .sphere = s }),
        Intersection.init(-4, Shape{ .sphere = s }),
    );
    try std.testing.expect(Intersections.approxEq(raySphere(r, s), expected));
}

test {
    _ = Intersections;
}
