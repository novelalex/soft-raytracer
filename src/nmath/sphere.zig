const v = @import("vector.zig");
const mathf = @import("mathf.zig");

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

    pub fn approxEq(self: Sphere, other: Sphere) bool {
        return mathf.approxEq(self.radius, other.radius) and v.approxEq(self.center, other.center);
    }
};
