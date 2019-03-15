from rest_framework import serializers

from api.models import DataPointsPh, DataPointsOxygen, DataPointsTemperature

class DataPointsPhSerializer(serializers.ModelSerializer):
    class Meta:
        model = DataPointsPh
        fields = ('id', 'datetime', 'value', 'systemId')

class DataPointsOxygenSerializer(serializers.ModelSerializer):
    class Meta:
        model = DataPointsOxygen
        fields = ('id', 'datetime', 'value', 'systemId')

class DataPointsTemperatureSerializer(serializers.ModelSerializer):
    class Meta:
        model = DataPointsTemperature
        fields = ('id', 'datetime', 'value', 'systemId')