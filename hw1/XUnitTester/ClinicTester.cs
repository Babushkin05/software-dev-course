namespace XUnitTester;
using Xunit;
using Moq;
using zoo_erp;

public class ClinicTester
{
    [Theory]
    [InlineData(100, true)]
    [InlineData(0, false)]
    [InlineData(50, false)]
    [InlineData(51, true)]
    /// <summary>
    /// Tests currectness of health detecting.
    /// </summary>
    /// <param name="health">Helth to compare</param>
    /// <param name="sentence">Right answer</param>
    public void Test_HealthChecking(byte health, bool sentence)
    {
        // Arrange
        Clinic clinic = new Clinic();
        var mockAnimal = new Mock<Animal>();
        mockAnimal.Setup(a => a.Health).Returns(health);

        // Act
        bool clinicSentence = clinic.IsHealthy(mockAnimal.Object);

        // Assert
        Assert.Equal(sentence, clinicSentence);
    }
}