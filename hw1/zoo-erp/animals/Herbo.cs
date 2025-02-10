using System.Reflection.Metadata;

namespace zoo_erp
{
    public abstract class Herbo : Animal
    {
        private const int INTERACT_THRESHOLD = 5;

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

        public bool IsInterqactive(){
            return Kindness > INTERACT_THRESHOLD;
        }

        public override string ToString()
        {
            return base.ToString() + $", kindness is {Kindness}";
        }
    }
}