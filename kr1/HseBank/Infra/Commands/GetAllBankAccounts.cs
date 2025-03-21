using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class GetAllBankAccountsCommand : ICommand
{
    private readonly IBankAccountService _bankAccountService;

    public GetAllBankAccountsCommand(IBankAccountService bankAccountService)
    {
        _bankAccountService = bankAccountService;
    }

    public object Execute()
    {
        return _bankAccountService.GetAllBankAccounts();
    }
}