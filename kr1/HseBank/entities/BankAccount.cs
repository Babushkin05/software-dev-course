using System.Data.Common;

namespace HseBank.entities
{
    public class BankAccount
    {
        public readonly uint id;

        public readonly string name;

        public int Balance
        {
            get;
            internal set;
        }

        internal BankAccount(uint id_, string name_ = "main account")
        {
            id = id_;
            name = name_;
            Balance = 0;
        }
    }
}