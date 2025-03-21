using HseBank.Domain;
using HseBank.UseCases;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace HseBank.Infra;

public class YamlBankAccountImporter : BankAccountImporter
{
    protected override List<BankAccount> ParseData(string fileContent)
{
    try
    {
        var deserializer = new DeserializerBuilder().Build();
        var accounts = deserializer.Deserialize<List<BankAccount>>(fileContent);
        return accounts ?? new List<BankAccount>();
    }
    catch
    {
        return new List<BankAccount>(); 
    }
}
}