using System;
namespace sem1
{
	internal class Engine
	{
		internal Engine(int size = 5)
		{
			Size = size;
		}

		private int _size;

		public int Size
		{
			get => _size;
			private set => _size = value;
		}
	}
}

