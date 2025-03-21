using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class GroupOperationsByCategoryCommand : ICommand
{
    private readonly IAnalyticsService _analyticsService;
    private readonly DateTime _startDate;
    private readonly DateTime _endDate;

    public GroupOperationsByCategoryCommand(IAnalyticsService analyticsService, DateTime startDate, DateTime endDate)
    {
        _analyticsService = analyticsService;
        _startDate = startDate;
        _endDate = endDate;
    }

    public object Execute()
    {
        return _analyticsService.GroupOperationsByCategory(_startDate, _endDate);
    }
}