using HseBank.Domain;
using HseBank.UseCases;
using Newtonsoft.Json;

namespace HseBank.Infra;

public class JsonBankAccountImporter : BankAccountImporter
{
    protected override List<BankAccount> ParseData(string fileContent)
    {
        return JsonConvert.DeserializeObject<List<BankAccount>>(fileContent)!;
    }
}