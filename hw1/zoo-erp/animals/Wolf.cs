namespace zoo_erp
{
    public class Wolf : Predator
    {
        public Wolf(string name, int food, byte health)
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