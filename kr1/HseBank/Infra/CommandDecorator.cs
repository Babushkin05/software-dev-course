using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

public class TimedCommandDecorator : ICommand
{
    private readonly ICommand _command;

    public TimedCommandDecorator(ICommand command)
    {
        _command = command;
    }

    public void Execute()
    {
        var stopwatch = System.Diagnostics.Stopwatch.StartNew();
        _command.Execute();
        stopwatch.Stop();
        Console.WriteLine($"Command executed in {stopwatch.ElapsedMilliseconds} ms.");
    }
}