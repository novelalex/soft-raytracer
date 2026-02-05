const std = @import("std");
const Color = @import("../nmath/color.zig").Color;

pub const Canvas = struct {
    width: usize,
    height: usize,
    buffer: []Color,

    pub fn init(allocator: std.mem.Allocator, width: usize, height: usize) !Canvas {
        const buffer_size = width * height;
        const buffer = try allocator.alloc(Color, buffer_size);

        for (buffer) |*p| {
            p.* = Color.init(0, 0, 0);
        }

        return .{
            .width = width,
            .height = height,
            .buffer = buffer,
        };
    }

    pub fn deinit(self: *Canvas, allocator: std.mem.Allocator) void {
        allocator.free(self.buffer);
    }

    pub fn pixelAt(self: Canvas, x: usize, y: usize) Color {
        return self.buffer[y * self.width + x];
    }

    pub fn writePixel(self: *Canvas, x: usize, y: usize, color: Color) void {
        self.buffer[y * self.width + x] = color;
    }

    fn constructPPMHeader(self: Canvas, allocator: std.mem.Allocator) ![]u8 {
        return try std.fmt.allocPrint(allocator, "P3\n{d} {d}\n255", .{ self.width, self.height });
    }

    fn constructPPMBody(self: Canvas, allocator: std.mem.Allocator) ![]u8 {
        var ppm_body: std.ArrayList(u8) = std.ArrayList(u8).init(allocator);
        errdefer ppm_body.deinit();

        // TODO:(Novel) print ppm body

        // This should reset on new line.
        var width_counter = 0;
        for (self.buffer) |c| {
            const r = std.math.clamp(std.math.round(c.rgba.r), 0, 255);
            var buf: [4]u8 = undefined;
            std.fmt.bufPrint(buf, "{d}", .{r});

            const g = std.math.clamp(std.math.round(c.rgba.g), 0, 255);
            const b = std.math.clamp(std.math.round(c.rgba.b), 0, 255);
            var buf: [32]u8 = undefined;

            ppm_body.print(allocator, "{} {d} {d}", .{ r, g, b });
        }

        return ppm_body.toOwnedSlice();
    }

    pub fn toPPM(self: Canvas, allocator: std.mem.Allocator) ![]u8 {
        const ppm_header = try self.constructPPMHeader(allocator);
        const ppm_body = try self.constructPPMBody(allocator);
        return try std.fmt.allocPrint(allocator, "{s}\n{s}\n", .{ ppm_header, ppm_body });
    }
};
