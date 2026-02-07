const vec = @import("vector.zig");
const matrix = @import("matrix.zig");
const std = @import("std");

pub const Quaternion = struct {
    w: f32,
    ijk: @Vector(3, f32),

    pub fn init(w: f32, ijk: @Vector(3, f32)) Quaternion {
        return .{
            .w = w,
            .ijk = ijk,
        };
    }

    pub fn multiply(lhs: Quaternion, rhs: Quaternion) Quaternion {
        return .{
            .w = (lhs.w * rhs.w) - @reduce(.Add, -(lhs.ikj * rhs.ijk)),
            .ijk = .{
                (lhs.w * rhs.ijk[0]) + (lhs.ijk[0] * rhs.w) - (lhs.ijk[2] * rhs.ijk[1]) + (lhs.ijk[1] * rhs.ijk[2]),
                (lhs.w * rhs.ijk[1]) + (lhs.ijk[1] * rhs.w) - (lhs.ijk[0] * rhs.ijk[2]) + (lhs.ijk[2] * rhs.ijk[0]),
                (lhs.w * rhs.ijk[2]) + (lhs.ijk[2] * rhs.w) - (lhs.ijk[1] * rhs.ijk[0]) + (lhs.ijk[0] * rhs.ijk[1]),
            },
        };
    }

    pub fn multiplyScalar(lhs: Quaternion, rhs: f32) Quaternion {
        return .{
            .w = lhs.w * rhs,
            .ijk = lhs.ijk * @as(@Vector(3, f32), @splat(rhs)),
        };
    }

    pub fn divideScalar(lhs: Quaternion, rhs: f32) Quaternion {
        return .{
            .w = lhs.w / rhs,
            .ijk = lhs.ijk / @as(@Vector(3, f32), @splat(rhs)),
        };
    }

    pub fn magnitude(q: Quaternion) f32 {
        return @sqrt((q.w * q.w) + @reduce(.Add, (q.ijk * q.ijk)));
    }

    pub fn conjugate(q: Quaternion) Quaternion {
        return .{
            .w = q.w,
            .ijk = -q.ijk,
        };
    }

    pub fn inverse(q: Quaternion) Quaternion {
        const mag = q.magnitude();
        const conj = q.conjugate();
        return conj.divideScalar(mag * mag);
    }

    pub fn normalize(q: Quaternion) Quaternion {
        return q.divideScalar(q.magnitude());
    }

    pub fn angleAxisRotation(theta: f32, axis: vec.Vec4) Quaternion {
        const rotation_axis = vec.vec3(vec.normalize(axis));
        const cos_value = std.math.cos(theta / 2.0);
        const sin_value = std.math.sin(theta / 2.0);
        const result = Quaternion.init(cos_value, rotation_axis * @as(vec.Vec3, @splat(sin_value)));
        return result.normalize();
    }

    pub fn toMatrix(q: Quaternion) matrix.Mat4 {
        // return matrix.mat4Transpose(matrix.Mat4{
        //     (1.0 - 2.0 * q.ijk[1] * q.ijk[1] - 2.0 * q.ijk[2] * q.ijk[2]), (2.0 * q.ijk[0] * q.ijk[1] + 2.0 * q.ijk[2] * q.w),            (2.0 * q.ijk[0] * q.ijk[2] - 2.0 * q.ijk[1] * q.w),            0.0,
        //     (2.0 * q.ijk[0] * q.ijk[1] - 2.0 * q.ijk[2] * q.w),            (1.0 - 2.0 * q.ijk[0] * q.ijk[0] - 2.0 * q.ijk[2] * q.ijk[2]), (2.0 * q.ijk[1] * q.ijk[2] + 2.0 * q.ijk[0] * q.w),            0.0,
        //     (2.0 * q.ijk[0] * q.ijk[2] + 2.0 * q.ijk[1] * q.w),            (2.0 * q.ijk[1] * q.ijk[2] - 2 * q.ijk[0] * q.w),              (1.0 - 2.0 * q.ijk[0] * q.ijk[0] - 2.0 * q.ijk[1] * q.ijk[1]), 0.0,
        //     0.0,                                                           0.0,                                                           0.0,                                                           1.0,
        // });

        return matrix.Mat4{
            (1.0 - 2.0 * q.ijk[1] * q.ijk[1] - 2.0 * q.ijk[2] * q.ijk[2]), (2.0 * q.ijk[0] * q.ijk[1] - 2.0 * q.ijk[2] * q.w),            (2.0 * q.ijk[0] * q.ijk[2] + 2.0 * q.ijk[1] * q.w),            0.0,
            (2.0 * q.ijk[0] * q.ijk[1] + 2.0 * q.ijk[2] * q.w),            (1.0 - 2.0 * q.ijk[0] * q.ijk[0] - 2.0 * q.ijk[2] * q.ijk[2]), (2.0 * q.ijk[1] * q.ijk[2] - 2 * q.ijk[0] * q.w),              0.0,
            (2.0 * q.ijk[0] * q.ijk[2] - 2.0 * q.ijk[1] * q.w),            (2.0 * q.ijk[1] * q.ijk[2] + 2.0 * q.ijk[0] * q.w),            (1.0 - 2.0 * q.ijk[0] * q.ijk[0] - 2.0 * q.ijk[1] * q.ijk[1]), 0.0,
            0.0,                                                           0.0,                                                           0.0,                                                           1.0,
        };
    }
};
