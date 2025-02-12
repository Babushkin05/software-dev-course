namespace zoo_erp
{
    public class Monkey : Herbo
    {
        public Monkey(string name, int food, byte health, byte kindness)
        {
            this.Name = name;
            this.Food = food;
            this.Kindness = kindness;
            this.Health = health;
        }

        public override string ToString()
        {
            return "Monkey " + base.ToString();
        }
    }
}