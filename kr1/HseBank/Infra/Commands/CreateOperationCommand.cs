using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;
public class CreateOperationCommand : ICommand
{
    private readonly IOperationService _operationService;
    private readonly Operation _operation;

    public CreateOperationCommand(IOperationService operationService, Operation operation)
    {
        _operationService = operationService;
        _operation = operation;
    }

    public object Execute()
    {
        _operationService.CreateOperation(_operation);
        return new object();
    }
}