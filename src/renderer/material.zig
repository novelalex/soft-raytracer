const clr = @import("../nmath/color.zig");
const v = @import("../nmath/vector.zig");
const PointLight = @import("light.zig").PointLight;
const std = @import("std");
const expect = std.testing.expect;
const constants = @import("../nmath/constants.zig");

pub const Material = struct {
    color: clr.Color = clr.init(1, 1, 1),
    ambient: f32 = 0.1,
    diffuse: f32 = 0.9,
    specular: f32 = 0.9,
    shininess: f32 = 200.0,

    pub fn init(color: clr.Color, ambient: f32, diffuse: f32, specular: f32, shininess: f32) Material {
        return .{
            .color = color,
            .ambient = ambient,
            .diffuse = diffuse,
            .specular = specular,
            .shininess = shininess,
        };
    }

    pub fn lighting(m: Material, l: PointLight, p: v.Point, eye: v.Vector, norm: v.Vector) clr.Color {
        var ambient: clr.Color = clr.init(0, 0, 0);
        var diffuse: clr.Color = clr.init(0, 0, 0);
        var specular: clr.Color = clr.init(0, 0, 0);

        const effective_color = m.color * l.intensity;
        const lightv = v.normalize(l.position - p);
        ambient = effective_color * v.splat(m.ambient);

        const light_dot_normal = v.dot(lightv, norm);
        if (light_dot_normal < 0) {
            diffuse = clr.init(0, 0, 0);
            specular = clr.init(0, 0, 0);
        } else {
            diffuse = effective_color * v.splat(m.diffuse * light_dot_normal);
            const reflectv = v.reflect(-lightv, norm);
            const reflect_dot_eye = v.dot(reflectv, eye);
            if (reflect_dot_eye <= 0) {
                specular = clr.init(0, 0, 0);
            } else {
                const factor = std.math.pow(f32, reflect_dot_eye, m.shininess);
                specular = l.intensity * v.splat(m.specular * factor);
            }
        }
        return clr.fromVec(ambient + diffuse + specular);
    }

    test "Material:lighting with eye between light and surface" {
        const m = Material{};
        const p = v.point(0, 0, 0);
        const eye = v.vector(0, 0, -1);
        const normal = v.vector(0, 0, -1);
        const light = PointLight.init(v.point(0, 0, -10), clr.init(1, 1, 1));
        const result = m.lighting(light, p, eye, normal);
        const expected = clr.init(1.9, 1.9, 1.9);
        try expect(v.approxEq(result, expected));
    }

    test "Material:lighting with the eye between light and surface, eye offset 45 degrees" {
        const m = Material{};
        const p = v.point(0, 0, 0);
        const root_2_over_2 = @sqrt(2.0) / 2.0;
        const eye = v.vector(0, root_2_over_2, -root_2_over_2);
        const normal = v.vector(0, 0, -1);
        const light = PointLight.init(v.point(0, 0, -10), clr.init(1, 1, 1));
        const result = m.lighting(light, p, eye, normal);
        const expected = clr.init(1.0, 1.0, 1.0);
        try expect(v.approxEq(result, expected));
    }

    test "Material:lighting with eye opposite surface, light offset 45 degrees" {
        const m = Material{};
        const p = v.point(0, 0, 0);
        const eye = v.vector(0, 0, -1);
        const normal = v.vector(0, 0, -1);
        const light = PointLight.init(v.point(0, 10, -10), clr.init(1, 1, 1));
        const result = m.lighting(light, p, eye, normal);
        const expected = clr.init(0.7364, 0.7364, 0.7364);
        try expect(v.approxEq(result, expected));
    }

    test "Material:lighting with eye in the path of the reflection vector" {
        const m = Material{};
        const p = v.point(0, 0, 0);
        const root_2_over_2 = @sqrt(2.0) / 2.0;
        const eye = v.vector(0, -root_2_over_2, -root_2_over_2);
        const normal = v.vector(0, 0, -1);
        const light = PointLight.init(v.point(0, 10, -10), clr.init(1, 1, 1));
        const result = m.lighting(light, p, eye, normal);
        const expected = clr.init(
            0.1 + 0.9 * root_2_over_2 + 0.9,
            0.1 + 0.9 * root_2_over_2 + 0.9,
            0.1 + 0.9 * root_2_over_2 + 0.9,
        );
        // math.pow is not exact so lets loosen up the approximate equal
        try expect(std.math.approxEqAbs(f32, result[0], expected[0], constants.epsilon_loose));
        try expect(std.math.approxEqAbs(f32, result[1], expected[1], constants.epsilon_loose));
        try expect(std.math.approxEqAbs(f32, result[2], expected[2], constants.epsilon_loose));
        try expect(std.math.approxEqAbs(f32, result[3], expected[3], constants.epsilon_loose));
    }

    test "Material:lighting with the light behind the surface" {
        const m = Material{};
        const p = v.point(0, 0, 0);
        const eye = v.vector(0, 0, -1);
        const normal = v.vector(0, 0, -1);
        const light = PointLight.init(v.point(0, 0, 10), clr.init(1, 1, 1));
        const result = m.lighting(light, p, eye, normal);
        const expected = clr.init(0.1, 0.1, 0.1);
        try expect(v.approxEq(result, expected));
    }
};
