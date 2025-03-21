using HseBank.Domain;
using HseBank.UseCases;
using HseBank.Infra;
using Microsoft.Extensions.DependencyInjection;
using System.Data.Common;

namespace HseBank.Ui;

public partial class ConsoleApp {
    private string _logpath;

    public ConsoleApp(string logFilePath) {
        _logpath = logFilePath;
    }

    public void Run(){
        var bankAccountServices = new ServiceCollection().AddSingleton<IBankAccountService, BankAccountService>();
        bankAccountServices.AddSingleton<IBankAccountService>(provider =>
        {
            var realService = provider.GetRequiredService<BankAccountService>();
            return new BankAccountServiceProxy(realService, _logpath);
        });

        var CategoryServices = new ServiceCollection().AddSingleton<ICategoryService, CategoryService>();
        CategoryServices.AddSingleton<ICategoryService>(provider => 
        {
            var realService = provider.GetRequiredService<CategoryService>();
            return new CategoryServiceProxy(realService, _logpath);
        });

        var OperationServices = new ServiceCollection().AddSingleton<IOperationService, OperationService>();
        OperationServices.AddSingleton<IOperationService>(provider => {
            var realService = provider.GetRequiredService<OperationService>();
            return new OperationServiceProxy(realService, _logpath);
        });

        var AnalyticsServices = new ServiceCollection().AddSingleton<IAnalyticsService, AnalyticsService>();
        AnalyticsServices.AddSingleton<IAnalyticsService>(provider => {
            var realService = provider.GetRequiredService<AnalyticsService>();
            return new AnalyticServiceProxy(realService, _logpath);
        });


        while(true){
            int choosen = Slider("What you want?", new string[]{"Analitics", "Bank Account", "Category", "Operationj"});
        }
    }
}