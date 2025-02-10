namespace zoo_erp
{
    public class Rabbit : Herbo
    {
        Rabbit(string name, int food, byte kindness, byte health)
        {
            this.Name = name;
            this.Food = food;
            this.Kindness = kindness;
            this.Health = health;
        }

        public override string ToString()
        {
            return "Rabbit " + base.ToString();
        }
    }
}