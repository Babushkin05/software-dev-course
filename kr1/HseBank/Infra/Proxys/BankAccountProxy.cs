using HseBank.UseCases;
using HseBank.Domain;
using Newtonsoft.Json;

namespace HseBank.Infra;
public class BankAccountServiceProxy : IBankAccountService
{
    private readonly IBankAccountService _realService;
    private readonly string _logFilePath;

    public BankAccountServiceProxy(IBankAccountService realService, string logFilePath)
    {
        _realService = realService;
        _logFilePath = logFilePath;
    }

    private void Log(string message)
    {
        string logMessage = $"{DateTime.Now}: {message}";
        File.AppendAllText(_logFilePath, logMessage + Environment.NewLine);
    }

    public void CreateBankAccount(BankAccount account)
    {
        Log($"CreateBankAccount called with account: {JsonConvert.SerializeObject(account)}");
        _realService.CreateBankAccount(account);
    }

    public void DeleteBankAccount(int id)
    {
        Log($"DeleteBankAccount called with id: {id}");
        _realService.DeleteBankAccount(id);
    }

    public BankAccount GetBankAccount(int id)
    {
        Log($"GetBankAccount called with id: {id}");
        return _realService.GetBankAccount(id);
    }

    public IEnumerable<BankAccount> GetAllBankAccounts()
    {
        Log("GetAllBankAccounts called");
        return _realService.GetAllBankAccounts();
    }
}