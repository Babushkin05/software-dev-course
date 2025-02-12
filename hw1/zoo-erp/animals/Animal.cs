namespace zoo_erp
{
    /// <summary>
    /// Abstract class of all animals
    /// </summary>
    public abstract class Animal : IAlive
    {
        /// <summary>
        /// Name of animal
        /// </summary>
        /// <value></value>
        public string? Name { get; private protected set; }

        /// <summary>
        /// Health of animal
        /// </summary>
        /// <value></value>
        public virtual byte Health {get; init; }

        /// <summary>
        /// Food consumption of animal
        /// </summary>
        /// <value></value>
        public virtual int Food { get; init; }

        public override string ToString()
        {
            return $"{Name} need {Food} kg food every day, health is {Health}";
        }
    }
}