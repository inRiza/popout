import { Image } from 'expo-image';
import { Platform, StyleProp, StyleSheet, View, Text } from 'react-native';

import { HelloWave } from '@/components/hello-wave';
import ParallaxScrollView from '@/components/parallax-scroll-view';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
// import { Link } from 'expo-router';

export default function HomeScreen() {
  return (
    <View style={[styles.container, {
        flexDirection: 'column'
      }]}
    >
      <View style={styles.card}>
        <Text style={styles.cardTitle}>Hang with Popout</Text>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
  },
  card: {
    flexDirection: 'column',
    backgroundColor: 'white',
    borderRadius: 10,
    padding: 10,
    shadowColor: 'black',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3,
    elevation: 5,
  },
  cardTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 10,
    color: 'black'
  }
});
