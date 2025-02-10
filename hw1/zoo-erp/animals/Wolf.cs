namespace zoo_erp
{
    public class Wolf : Predator
    {
        Wolf(string name, int food, byte health)
        {
            this.Name = name;
            this.Food = food;
            this.Health = health;
        }

        public override string ToString()
        {
            return "Wolf " + base.ToString();
        }
    }
}