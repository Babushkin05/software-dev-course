namespace zoo_erp
{
    public abstract class Animal : IAlive
    {
        public uint AnimalId { get; private protected set; }

        public string? Name { get; private protected set; }

        public byte Health {get; private protected set; }

        public int Food { get; init; }
    }
}