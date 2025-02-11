namespace zoo_erp
{
    public abstract class Animal : IAlive
    {
        public string? Name { get; private protected set; }

        public byte Health {get; private protected set; }

        public int Food { get; init; }

        public override string ToString()
        {
            return $"{Name} need {Food} kg food every day, health is {Health}";
        }
    }
}