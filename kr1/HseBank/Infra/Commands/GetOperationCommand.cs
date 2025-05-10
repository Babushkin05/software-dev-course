using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;
public class GetOperationCommand : ICommand
{
    private readonly IOperationService _operationService;
    private readonly UInt128 _id;

    public GetOperationCommand(IOperationService operationService, UInt128 id)
    {
        _operationService = operationService;
        _id = id;
    }

    public object Execute()
    {
        return _operationService.GetOperation(_id);
    }
}