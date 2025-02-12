using System.Reflection.Metadata;

namespace zoo_erp
{
    /// <summary>
    /// Abstract class of all Herbos
    /// </summary>
    public abstract class Herbo : Animal
    {
        /// <summary>
        /// Threshold of interactive
        /// </summary>
        private const int INTERACT_THRESHOLD = 5;

        /// <summary>
        /// kindness of herbo
        /// </summary>
        private byte _kindness;
        public byte Kindness
        {
            get { return _kindness; }
            set
            {
                if (value > 10)
                {
                    throw new ArgumentException("Kindness value must be between 0 and 10.");
                }
                _kindness = value;
            }
        }

        /// <summary>
        /// Check kindness of animal
        /// </summary>
        /// <returns>Is people can interact with it</returns>
        public bool IsInterqactive()
        {
            return Kindness > INTERACT_THRESHOLD;
        }

        public override string ToString()
        {
            return base.ToString() + $", kindness is {Kindness}";
        }
    }
}