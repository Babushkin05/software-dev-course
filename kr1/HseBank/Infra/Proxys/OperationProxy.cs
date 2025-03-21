using HseBank.Domain;
using HseBank.UseCases;

namespace HseBank.Infra;

using System;
using System.Collections.Generic;
using System.IO;

public class OperationServiceProxy : IOperationService
{
    private readonly IOperationService _realService;
    private readonly string _logFilePath;

    public OperationServiceProxy(IOperationService realService, string logFilePath)
    {
        _realService = realService;
        _logFilePath = logFilePath;
    }

    private void Log(string message)
    {
        string logMessage = $"{DateTime.Now}: {message}";
        File.AppendAllText(_logFilePath, logMessage + Environment.NewLine);
    }

    public void CreateOperation(Operation operation)
    {
        Log($"Creating operation with ID {operation.Id}");
        try
        {
            _realService.CreateOperation(operation);
            Log($"Operation with ID {operation.Id} created successfully.");
        }
        catch (Exception ex)
        {
            Log($"Error creating operation: {ex.Message}");
            throw;
        }
    }

    public void DeleteOperation(UInt128 id)
    {
        Log($"Deleting operation with ID {id}");
        try
        {
            _realService.DeleteOperation(id);
            Log($"Operation with ID {id} deleted successfully.");
        }
        catch (Exception ex)
        {
            Log($"Error deleting operation: {ex.Message}");
            throw;
        }
    }

    public Operation GetOperation(UInt128 id)
    {
        Log($"Getting operation with ID {id}");
        try
        {
            var operation = _realService.GetOperation(id);
            Log($"Retrieved operation with ID {id}");
            return operation;
        }
        catch (Exception ex)
        {
            Log($"Error getting operation: {ex.Message}");
            throw;
        }
    }

    public IEnumerable<Operation> GetAllOperations()
    {
        Log("Getting all operations");
        try
        {
            var operations = _realService.GetAllOperations();
            Log($"Retrieved {operations.Count()} operations.");
            return operations;
        }
        catch (Exception ex)
        {
            Log($"Error getting all operations: {ex.Message}");
            throw;
        }
    }
}