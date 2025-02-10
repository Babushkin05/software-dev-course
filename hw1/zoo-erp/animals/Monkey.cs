namespace zoo_erp
{
    public class Monkey : Herbo
    {
        Monkey(string name, int food, byte kindness, byte health)
        {
            this.Name = name;
            this.Food = food;
            this.Kindness = kindness;
            this.Health = health;
        }
    }
}