using Context;
using Dapr.Client;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddDaprClient();
builder.Services.AddDbContext<ApplicationDbContext>(
    o => o.UseNpgsql(builder.Configuration["ConnectionString"]!)
);

var app = builder.Build();

app.UseRouting();

app.MapGet("/do-work", async (DaprClient client) =>
{
    Console.WriteLine("Working on something");

    // Defer event save
    await client.PublishEventAsync(
        "pubsub",
        "event.create",
        new Event
        {
            Data = $"Testing data {DateTime.UtcNow}"
        }
    );
});

app.MapGet("/events", async (ApplicationDbContext context) =>
{
    Console.WriteLine("Fetching events for UI");
    return await context.Events.ToListAsync();
});

app.Run();
