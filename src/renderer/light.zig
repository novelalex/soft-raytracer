const clr = @import("../nmath/color.zig");
const vec = @import("../nmath/vector.zig");
const std = @import("std");
const expect = std.testing.expect;

pub const PointLight = struct {
    position: vec.Point,
    intensity: clr.Color,

    pub fn init(position: vec.Point, intensity: clr.Color) PointLight {
        return .{
            .position = position,
            .intensity = intensity,
        };
    }
};
