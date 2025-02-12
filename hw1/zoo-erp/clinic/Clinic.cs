using System.Reflection.Metadata;

namespace zoo_erp
{
    public class Clinic : IClinic
    {
        /// <summary>
        /// Threshold of healty
        /// </summary>
        private const byte HEALTH_THRESHOLD = 50;

        public bool IsHealthy(Animal animal){
            return animal.Health > HEALTH_THRESHOLD;
        }
    }
}