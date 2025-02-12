using Microsoft.Extensions.DependencyInjection;
namespace zoo_erp
{
    /// <summary>
    /// Main class
    /// </summary>
    internal class Program
    {
        /// <summary>
        /// Main Program
        /// </summary>
        internal static void Main(){
            // DI Container of Clinic
            var services = new ServiceCollection().AddTransient<IClinic, Clinic>();
            services.AddTransient<Zoo>();

            using ServiceProvider serviceProvider = services.BuildServiceProvider();
            var zoo = serviceProvider.GetService<Zoo>();

            // Initing console app
            ConsoleApp consoleApp = new ConsoleApp(zoo);

            // Run program
            consoleApp.Run();

        }
    }
}
