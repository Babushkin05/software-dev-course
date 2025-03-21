using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class AnalyticServiceProxy : IAnalyticsService
{
    private readonly IAnalyticsService _realService;
    private readonly string _logFilePath;

    public AnalyticServiceProxy(IAnalyticsService realService, string logFilePath)
    {
        _realService = realService;
        _logFilePath = logFilePath;
    }

    private void Log(string message)
    {
        string logMessage = $"{DateTime.Now}: {message}";
        File.AppendAllText(_logFilePath, logMessage + Environment.NewLine);
    }

    public decimal CalculateBalanceDifference(DateTime startDate, DateTime endDate) {
        decimal bal = _realService.CalculateBalanceDifference(startDate, endDate);
        Log($"CalculateBalanceDifference called: {bal}");
        return bal;
    }

    public Dictionary<string, long> GroupOperationsByCategory(DateTime startDate, DateTime endDate){
        Log("GroupOperationsByCategory called");
        return _realService.GroupOperationsByCategory(startDate, endDate);
    }
}