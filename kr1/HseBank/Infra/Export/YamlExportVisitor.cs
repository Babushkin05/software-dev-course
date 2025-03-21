using HseBank.Domain;
using HseBank.UseCases;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace HseBank.Infra;


public class YamlExportVisitor : IVisitor
{
    private readonly List<BankAccount> _accounts = new();

    public void Visit(BankAccount account)
    {
        _accounts.Add(account);
    }

    public void SaveToFile(string filePath)
    {
        var serializer = new SerializerBuilder().Build();
        string yaml = serializer.Serialize(_accounts);
        File.WriteAllText(filePath, yaml);
    }
}