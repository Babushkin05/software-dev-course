using HseBank.UseCases;
namespace HseBank.Infra;

public interface ICommand {
    object Execute();
}