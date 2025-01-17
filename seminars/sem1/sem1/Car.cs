using System;
using System.Drawing;

namespace sem1
{
	public class Car
	{
        public int Number { get; private set; }

		public Engine Engine { get; private set; }

		public Car()
		{
			++cars_number;
			Number = cars_number;
			Engine = new Engine { Size = _random.Next(1, 10) };
		}

        private static readonly Random _random = new();

        private static int cars_number = 0;

        public override string ToString()
        {
			return $"Car number {Number}, Engine size {Engine.Size}";
        }
    }
}

