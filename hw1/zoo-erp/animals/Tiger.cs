namespace zoo_erp
{
    public class Tiger : Predator
    {
        Tiger(string name, int food, byte health)
        {
            this.Name = name;
            this.Food = food;
            this.Health = health;
        }

        public override string ToString()
        {
            return "Tiger " + base.ToString();
        }
    }
}