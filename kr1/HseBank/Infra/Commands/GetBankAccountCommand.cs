using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class GetBankAccountCommand : ICommand
{
    private readonly IBankAccountService _bankAccountService;
    private readonly int _accountId;

    public GetBankAccountCommand(IBankAccountService bankAccountService, int accountId)
    {
        _bankAccountService = bankAccountService;
        _accountId = accountId;
    }

    public object Execute()
    {
        return _bankAccountService.GetBankAccount(_accountId);
    }
}