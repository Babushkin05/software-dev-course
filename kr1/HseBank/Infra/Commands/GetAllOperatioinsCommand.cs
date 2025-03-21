using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;
public class GetAllOperationsCommand : ICommand
{
    private readonly IOperationService _operationService;

    public GetAllOperationsCommand(IOperationService operationService)
    {
        _operationService = operationService;
    }

    public object Execute()
    {
        return _operationService.GetAllOperations();
    }
}