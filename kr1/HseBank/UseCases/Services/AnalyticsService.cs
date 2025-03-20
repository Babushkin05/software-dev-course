using HseBank.Domain;
namespace HseBank.UseCases;

public class AnalyticsService
{
    private readonly OperationService _operationService;

    public AnalyticsService(OperationService operationService)
    {
        _operationService = operationService;
    }

    public decimal CalculateBalanceDifference(DateTime startDate, DateTime endDate)
    {
        var operations = _operationService.GetAllOperations()
            .Where(o => o.Time >= startDate && o.Time <= endDate);

        decimal totalIncome = operations.Where(o => o.Type == OperationType.Income).Sum(o => o.Amount);
        decimal totalExpense = operations.Where(o => o.Type == OperationType.Consumption).Sum(o => o.Amount);

        return totalIncome - totalExpense;
    }

    public Dictionary<string, long> GroupOperationsByCategory(DateTime startDate, DateTime endDate)
    {
        var operations = _operationService.GetAllOperations()
            .Where(o => o.Time >= startDate && o.Time <= endDate);

        return operations
            .GroupBy(o => o.CategoryId)
            .ToDictionary(
                g => g.First().CategoryId.ToString(),
                g => g.Sum(o => o.Amount)
            );
    }
}