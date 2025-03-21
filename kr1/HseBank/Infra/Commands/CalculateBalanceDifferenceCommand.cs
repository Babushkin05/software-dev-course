using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class CalculateBalanceDifferenceCommand : ICommand
{
    private readonly IAnalyticsService _analyticsService;
    private readonly DateTime _startDate;
    private readonly DateTime _endDate;

    public CalculateBalanceDifferenceCommand(IAnalyticsService analyticsService, DateTime startDate, DateTime endDate)
    {
        _analyticsService = analyticsService;
        _startDate = startDate;
        _endDate = endDate;
    }

    public object Execute()
    {
        return _analyticsService.CalculateBalanceDifference(_startDate, _endDate);
    }
}