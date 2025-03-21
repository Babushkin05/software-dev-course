using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class CsvBankAccountImporter : BankAccountImporter
{
    protected override List<BankAccount> ParseData(string fileContent)
    {
        var accounts = new List<BankAccount>();
        var lines = fileContent.Split('\n');

        foreach (var line in lines)
        {
            var parts = line.Split(',');
            if (parts.Length >= 2 && uint.TryParse(parts[0], out uint id))
            {
                var account = new BankAccount(id, parts[1]);
                accounts.Add(account);
            }
        }

        return accounts;
    }
}