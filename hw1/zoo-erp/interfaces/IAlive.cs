namespace zoo_erp
{
    /// <summary>
    /// Alive interface
    /// </summary>
    internal interface IAlive
    {
        /// <summary>
        /// Kg of food that aliver need every day
        /// </summary>
        /// <value></value>
        int Food { get; init; }

        /// <summary>
        /// Health of aliver
        /// </summary>
        /// <value></value>
        byte Health {get; init; }
    }
}