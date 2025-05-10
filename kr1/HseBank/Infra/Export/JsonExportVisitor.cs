using HseBank.Domain;
using HseBank.UseCases;
using Newtonsoft.Json;

namespace HseBank.Infra;

public class JsonExportVisitor : IVisitor
{
    private readonly List<BankAccount> _accounts = new();

    public void Visit(BankAccount account)
    {
        _accounts.Add(account);
    }

    public void SaveToFile(string filePath)
    {
        string json = JsonConvert.SerializeObject(_accounts, Formatting.Indented);
        File.WriteAllText(filePath, json);
    }
}