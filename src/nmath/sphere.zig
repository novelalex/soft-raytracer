const v = @import("vector.zig");
const matrix = @import("matrix.zig");
const mathf = @import("mathf.zig");
const std = @import("std");

pub const Sphere = struct {
    center: v.Point,
    radius: f32,

    pub fn init(center: v.Point, radius: f32) Sphere {
        return .{
            .center = center,
            .radius = radius,
        };
    }

    pub const unit = Sphere{
        .center = v.point(0, 0, 0),
        .radius = 1,
    };

    pub fn getTransform(s: Sphere) matrix.Mat4 {
        const TS = matrix.Mat4{
            s.radius, 0,        0,        v.x(s.center),
            0,        s.radius, 0,        v.y(s.center),
            0,        0,        s.radius, v.y(s.center),
            0,        0,        0,        1,
        };

        // TODO: (Novel) Add rotation, we dont really need it
        // at the moment but it will be handy when we add textures.
        return TS;
    }

    test "Sphere:getTransform of a unit sphere is the identity matrix" {
        const s = Sphere.unit;
        const expected = matrix.mat4_identity;
        try std.testing.expect(matrix.approxEq(matrix.Mat4, s.getTransform(), expected));
    }

    pub fn approxEq(self: Sphere, other: Sphere) bool {
        return mathf.approxEq(self.radius, other.radius) and v.approxEq(self.center, other.center);
    }
};

test {
    _ = Sphere;
}
