using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using Google.Protobuf.WellKnownTypes;
using Microsoft.EntityFrameworkCore;

namespace Context;

public class ApplicationDbContext : DbContext
{
    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options) : base(options) { }

    public DbSet<Event> Events { get; set; } = default!;
}

[Table(nameof(Event))]
public class Event
{
    [Key]
    public int Id { get; set; }
    public string? Data { get; set; }
    public DateTime CreatedAt { get; private set; } = DateTime.UtcNow;
}