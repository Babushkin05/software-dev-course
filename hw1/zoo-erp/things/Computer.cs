namespace zoo_erp
{
    public class Computer : Thing
    {
        public string? computerInfo;

        public Computer(string? computer)
        {
            computerInfo = computer;
        }

        public override string ToString()
        {
            return base.ToString() + $" Computer {computerInfo}";
        }
    }
}