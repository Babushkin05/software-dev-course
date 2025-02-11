namespace zoo_erp
{
    public class Computer : Thing
    {
        public string? computerInfo;

        Computer(string? computer)
        {
            computerInfo = computer;
        }

        public override string ToString()
        {
            return base.ToString() + $" Computer {computerInfo}";
        }
    }
}