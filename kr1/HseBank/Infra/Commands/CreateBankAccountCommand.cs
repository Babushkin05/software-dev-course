using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class CreateBankAccountCommand : ICommand
{
    private readonly IBankAccountService _bankAccountService;
    private readonly BankAccount _bankAccount;

    public CreateBankAccountCommand(IBankAccountService bankAccountService, BankAccount bankAccount)
    {
        _bankAccountService = bankAccountService;
        _bankAccount = bankAccount;
    }

    public object Execute()
    {
        _bankAccountService.CreateBankAccount(_bankAccount);
        return new object();
    }
}