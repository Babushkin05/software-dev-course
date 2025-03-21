using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;
public class DeleteOperationCommand : ICommand
{
    private readonly IOperationService _operationService;
    private readonly UInt128 _id;

    public DeleteOperationCommand(IOperationService operationService, UInt128 id)
    {
        _operationService = operationService;
        _id = id;
    }

    public object Execute()
    {
        _operationService.DeleteOperation(_id);
        return new object();
    }
}