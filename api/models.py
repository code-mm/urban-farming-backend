from django.db import models

class System(models.Model):
    systemId = models.CharField(max_length=5)
    systemSecret = models.CharField(max_length=20)

class DataPointsPh(models.Model):
    datetime = models.DateTimeField()
    value = models.FloatField()
    systemId = models.ForeignKey(System, on_delete=models.CASCADE)

class DataPointsOxygen(models.Model):
    datetime = models.DateTimeField()
    value = models.FloatField()
    systemId = models.ForeignKey(System, on_delete=models.CASCADE)

class DataPointsTemperature(models.Model):
    datetime = models.DateTimeField()
    value = models.FloatField()
    systemId = models.ForeignKey(System, on_delete=models.CASCADE)