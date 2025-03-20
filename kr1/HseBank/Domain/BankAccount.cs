using System.Data.Common;

namespace HseBank.Domain
{
    public class BankAccount
    {
        private List<Operation> operationsStory;

        public readonly uint Id;

        public readonly string Name;

        public int Balance
        {
            get;
            private set;
        }

        internal BankAccount(uint id_, string name_ = "main account")
        {
            Id = id_;
            Name = name_;
            Balance = 0;
        }

        public List<Operation> GetHistory(){
            return operationsStory;
        }

        public void AddOperation(Operation operation){
            if(operation.Type == OperationType.Income){
                Balance += (int)operation.Amount;
            } 
            else{
                Balance -= (int)operation.Amount;
            }
            operationsStory.Add(operation);
        }
    }
}