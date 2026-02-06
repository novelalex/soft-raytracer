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
        var ppm_body: std.ArrayList(u8) = .empty;
        errdefer ppm_body.deinit(allocator);

        var width_counter: usize = 0;
        var digit_buffer: [4]u8 = undefined;

        var at_line_start = true;
        for (self.buffer) |c| {
            inline for (0..3) |j| {
                const color_channel: u8 = @intFromFloat(
                    std.math.clamp(std.math.round(c.vec.at(j) * 255), 0, 255),
                );
                const written_slice =
                    try std.fmt.bufPrint(&digit_buffer, "{d}", .{color_channel});

                const width_needed = written_slice.len + @intFromBool(!at_line_start);

                if (width_counter + width_needed > 70) {
                    try ppm_body.print(allocator, "\n", .{});
                    width_counter = 0;
                    at_line_start = true;
                }

                if (!at_line_start) {
                    try ppm_body.print(allocator, " ", .{});
                }

                try ppm_body.print(allocator, "{s}", .{written_slice});
                width_counter += width_needed;
                at_line_start = false;
            }
        }

        try ppm_body.print(allocator, "\n\n", .{});

        return ppm_body.toOwnedSlice(allocator);
    }

    pub fn toPPM(self: Canvas, allocator: std.mem.Allocator) ![]u8 {
        const ppm_header = try self.constructPPMHeader(allocator);
        const ppm_body = try self.constructPPMBody(allocator);
        return try std.fmt.allocPrint(allocator, "{s}\n{s}\n", .{ ppm_header, ppm_body });
    }
};
