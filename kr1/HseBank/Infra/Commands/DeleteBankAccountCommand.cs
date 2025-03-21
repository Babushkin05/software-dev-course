using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class DeleteBankAccountCommand : ICommand
{
    private readonly IBankAccountService _bankAccountService;
    private readonly int _accountId;

    public DeleteBankAccountCommand(IBankAccountService bankAccountService, int accountId)
    {
        _bankAccountService = bankAccountService;
        _accountId = accountId;
    }

    public object Execute()
    {
        _bankAccountService.DeleteBankAccount(_accountId);
        return new object();
    }
}