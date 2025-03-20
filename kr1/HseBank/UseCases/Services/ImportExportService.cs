using HseBank.Domain;

public class BankAccountService
{
    private readonly List<BankAccount> _bankAccounts = new();

    public void CreateBankAccount(BankAccount account)
    {
        if (account == null)
            throw new ArgumentNullException(nameof(account));

        if (_bankAccounts.Any(a => a.Id == account.Id))
            throw new InvalidOperationException("Account with the same ID already exists.");

        _bankAccounts.Add(account);
    }

    public void DeleteBankAccount(int id)
    {
        var account = _bankAccounts.FirstOrDefault(a => a.Id == id);
        if (account == null)
            throw new InvalidOperationException("Account not found.");

        _bankAccounts.Remove(account);
    }

    public BankAccount? GetBankAccount(int id)
    {
        return _bankAccounts.FirstOrDefault(a => a.Id == id);
    }

    public IEnumerable<BankAccount> GetAllBankAccounts()
    {
        return _bankAccounts;
    }
}