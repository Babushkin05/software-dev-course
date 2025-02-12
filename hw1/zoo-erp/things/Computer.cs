namespace zoo_erp
{
    public class Computer : Thing
    {
        /// <summary>
        /// Data about computer
        /// </summary>
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