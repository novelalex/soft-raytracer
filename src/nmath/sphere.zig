const v = @import("vector.zig");
const matrix = @import("matrix.zig");
const mathf = @import("mathf.zig");
const std = @import("std");

pub const Sphere = struct {
    transform: matrix.Mat4,

    pub fn init() Sphere {
        return .{ .transform = matrix.mat4_identity };
    }

    test "Sphere:getTransform of a unit sphere is the identity matrix" {
        const s = Sphere.init();
        const expected = matrix.mat4_identity;
        try std.testing.expect(matrix.approxEq(matrix.Mat4, s.transform, expected));
    }

    pub fn approxEq(self: Sphere, other: Sphere) bool {
        return matrix.approxEq(matrix.Mat4, self.transform, other.transform);
    }
};

test {
    _ = Sphere;
}
