using System.Reflection.Metadata;

namespace zoo_erp
{
    internal class Clinic : IClinic
    {
        private const byte HEALTH_THRESHOLD = 50;

        public bool IsHealthy(Animal animal){
            return animal.Health > HEALTH_THRESHOLD;
        }
    }
}