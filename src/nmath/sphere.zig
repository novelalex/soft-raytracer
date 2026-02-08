const v = @import("vector.zig");
const matrix = @import("matrix.zig");
const transform = @import("transform.zig");
const mathf = @import("mathf.zig");
const std = @import("std");
const Material = @import("../renderer/material.zig").Material;

pub const Sphere = struct {
    transform: matrix.Mat4,
    material: Material = Material{},

    pub fn init() Sphere {
        return .{ .transform = matrix.mat4_identity };
    }

    pub fn withTransform(s: Sphere, t: matrix.Mat4) Sphere {
        return .{ .transform = matrix.mat4Multiply(s.transform, t) };
    }

    test "Sphere:getTransform of a unit sphere is the identity matrix" {
        const s = Sphere.init();
        const expected = matrix.mat4_identity;
        try std.testing.expect(matrix.approxEq(matrix.Mat4, s.transform, expected));
    }

    pub fn approxEq(self: Sphere, other: Sphere) bool {
        return matrix.approxEq(matrix.Mat4, self.transform, other.transform);
    }

    pub fn normalAt(sphere: Sphere, world_point: v.Point) v.Vector {
        const object_point = matrix.mat4MultiplyVec(matrix.mat4Inverse(sphere.transform), world_point);
        const object_normal = object_point - v.point(0, 0, 0);
        var world_normal = matrix.mat4MultiplyVec(matrix.mat4Transpose(matrix.mat4Inverse(sphere.transform)), object_normal);
        world_normal[3] = 0;
        return v.normalize(world_normal);
    }

    test "Sphere:normalAt works for unit spheres" {
        const s = Sphere.init();
        const n = s.normalAt(v.point(@sqrt(3.0) / 3.0, @sqrt(3.0) / 3.0, @sqrt(3.0) / 3.0));
        const expected = v.vector(@sqrt(3.0) / 3.0, @sqrt(3.0) / 3.0, @sqrt(3.0) / 3.0);
        try std.testing.expect(v.approxEq(n, expected));
        try std.testing.expect(v.approxEq(n, v.normalize(n)));
    }

    test "Sphere:normalAt works for translated spheres" {
        var s = Sphere.init()
            .withTransform(transform.translation(0, 1, 0));
        const n = s.normalAt(v.point(0, 1.70711, -0.70711));
        const expected = v.vector(0, 0.70711, -0.70711);
        try std.testing.expect(v.approxEq(n, expected));
        try std.testing.expect(v.approxEq(n, v.normalize(n)));
    }
};

test {
    _ = Sphere;
}
