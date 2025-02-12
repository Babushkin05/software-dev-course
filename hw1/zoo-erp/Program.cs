using Microsoft.Extensions.DependencyInjection;
namespace zoo_erp
{
    internal class Program
    {
        internal static void Main(){
            var services = new ServiceCollection().AddTransient<IClinic, Clinic>();
            services.AddTransient<Zoo>();
            using ServiceProvider serviceProvider = services.BuildServiceProvider();

            var zoo = serviceProvider.GetService<Zoo>();

            ConsoleApp consoleApp = new ConsoleApp(zoo);

            consoleApp.Run();

        }
    }
}
